{{/* vim: set filetype=mustache: */}}


{{/*
Common labels
*/}}
{{- define "example-backend.labels" -}}
helm.sh/chart: {{ include "example-backend.chart" . }}
{{ include "example-backend.selectorLabels" . }}
{{- if $.Chart.AppVersion }}
app.kubernetes.io/version: {{ $.Chart.AppVersion | quote }}
{{- end }}
app.kubernetes.io/managed-by: {{ .Release.Service }}
{{- end -}}

{{/*
Selector labels
*/}}
{{- define "example-backend.selectorLabels" -}}
app.kubernetes.io/name: {{ include "example-backend.name" . }}
app.kubernetes.io/instance: {{ $.Release.Name }}
app.kubernetes.io/component: {{ .Chart.Name }}
{{- end -}}

{{/*
Expand the name of the chart.
*/}}
{{- define "example-backend.name" -}}
{{- default .Chart.Name .Values.nameOverride | trunc 63 | trimSuffix "-" -}}
{{- end -}}

{{/*
Create a default fully qualified app name.
We truncate at 63 chars because some Kubernetes name fields are limited to this (by the DNS naming spec).
If release name contains chart name it will be used as a full name.
*/}}
{{- define "example-backend.fullname" -}}
{{- if $.Values.fullnameOverride -}}
{{- $.Values.fullnameOverride | trunc 63 | trimSuffix "-" -}}
{{- else -}}
{{- $name := default .Chart.Name $.Values.nameOverride -}}
{{- if contains $name .Chart.Name -}}
{{- .Chart.Name | trunc 63 | trimSuffix "-" -}}
{{- else -}}
{{- printf "%s-%s" .Chart.Name $name | trunc 63 | trimSuffix "-" -}}
{{- end -}}
{{- end -}}
{{- end -}}

{{/*
Create chart name and version as used by the chart label.
*/}}
{{- define "example-backend.chart" -}}
{{- printf "%s-%s" .Chart.Name .Chart.Version | replace "+" "_" | trunc 63 | trimSuffix "-" -}}
{{- end -}}

{{/*
Create tls secret name
*/}}
{{- define "example-backend.tlsSecretName" -}}
{{-  .Release.Name  -}}-tls
{{- end -}}

{{/*
Create the name of the service account to use
*/}}
{{- define "example-backend.serviceAccountName" -}}
{{- default .Chart.Name .Values.serviceAccount.name  -}}
{{- end -}}
