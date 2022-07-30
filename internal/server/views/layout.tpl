<!DOCTYPE html>
<html>
  <head>
    <title>Terralist | {{ .Values.Title | default "Home" }}</title>

    <style>
      {{ block "style" .}}
      {{ end }}
    </style>
  </head>

  <body>
    <div class="main">
      {{ block "content" . }}
        <!-- Fallback -->
        <h2>404 not found</h2>
      {{ end }}
    </div>

    <script type="text/javascript">
      {{ block "script" . }}
      {{ end }}
    </script>
  </body>
</html>