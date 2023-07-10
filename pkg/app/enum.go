package app

type TraceExporter string

const (
	TraceExporterJaeger     TraceExporter = "jaeger"
	TraceExporterPrometheus TraceExporter = "prometheus"
)
