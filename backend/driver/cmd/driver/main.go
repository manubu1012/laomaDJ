package main

import (
	"flag"
	"github.com/go-kratos/kratos/contrib/registry/consul/v2"
	"github.com/google/uuid"
	"github.com/hashicorp/consul/api"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/sdk/resource"
	traceSDK "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
	"os"

	"driver/internal/conf"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"

	_ "go.uber.org/automaxprocs"
)

// go build -ldflags "-X main.Version=x.y.z"
var (
	// Name is the name of the compiled software.
	Name = "Driver"
	// Version is the version of the compiled software.
	Version = "1.0.0"
	// flagconf is the config flag.
	flagconf string

	id = Name+uuid.NewString()
)

func init() {
	flag.StringVar(&flagconf, "conf", "../../configs", "config path, eg: -conf config.yaml")
}

func newApp(cs *conf.Service, logger log.Logger, gs *grpc.Server, hs *http.Server) *kratos.App {
	log.Info(cs.Consul.Addr)
	reg,err:=initConsul(cs.Consul.Addr)
	log.Info("----------------")
	if err!=nil{
		panic(err)
	}
	err=initTrace(cs.Jaeger.Url)
	if err!=nil{
		panic(err)
	}
	return kratos.New(
		kratos.ID(id),
		kratos.Name(Name),
		kratos.Version(Version),
		kratos.Metadata(map[string]string{}),
		kratos.Logger(logger),
		kratos.Server(
			gs,
			hs,
		),
		kratos.Registrar(reg),
	)
}

func main() {
	flag.Parse()
	logger := log.With(log.NewStdLogger(os.Stdout),
		"ts", log.DefaultTimestamp,
		"caller", log.DefaultCaller,
		"service.id", id,
		"service.name", Name,
		"service.version", Version,
		"trace.id", tracing.TraceID(),
		"span.id", tracing.SpanID(),
	)
	c := config.New(
		config.WithSource(
			file.NewSource(flagconf),
		),
	)
	defer c.Close()

	if err := c.Load(); err != nil {
		panic(err)
	}

	var bc conf.Bootstrap
	if err := c.Scan(&bc); err != nil {
		panic(err)
	}

	app, cleanup, err := wireApp(bc.Service,bc.Server, bc.Data, logger)
	if err != nil {
		panic(err)
	}
	defer cleanup()

	// start and wait for stop signal
	if err := app.Run(); err != nil {
		panic(err)
	}
}


func initConsul(addr string) (*consul.Registry ,error){
	consulConfig:=api.DefaultConfig()
	consulConfig.Address= addr
	client, err := api.NewClient(consulConfig)
	if err != nil {
		return nil,err
	}
	// new reg with consul client
	reg := consul.New(client)
	return reg,nil
}

func initTrace(url string)error	  {
	exporter, err := jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(url)))
	if err!=nil{
		return err
	}
	traceProvider:=traceSDK.NewTracerProvider(
		traceSDK.WithSampler(traceSDK.AlwaysSample()),
		traceSDK.WithBatcher(exporter),
		traceSDK.WithResource(resource.NewSchemaless(
			semconv.ServiceNameKey.String(Name),
			attribute.String("exporter","jaeger"),
		)),
	)
	otel.SetTracerProvider(traceProvider)
	return  nil
}