{{ define "style" }}
{{ end }}

{{ define "content" }}
<h2>Sign In with {{ .Values.Provider }}</h2>
<p>This resource is protected and requires authentication using your {{ .Values.Provider }} account.</p>
<form method="post">
  <button type="submit">Continue</button>
</form>
{{ end }}

{{ define "script" }}
{{ end }}