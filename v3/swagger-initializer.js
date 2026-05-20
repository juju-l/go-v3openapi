window.onload = function() {
  //<editor-fold desc="Changeable Configuration Block">

  // the following lines will be replaced by docker/configurator, when it runs in a docker-container
  window.ui = SwaggerUIBundle({
    url: "/openapi/v3",
    dom_id: '#swagger-ui',
    deepLinking: true,
    presets: [
      SwaggerUIBundle.presets.apis,
      SwaggerUIStandalonePreset
    ],
    plugins: [
      SwaggerUIBundle.
      plugins.
      DownloadUrl
    ],
    layout: "StandaloneLayout",
    docExpansion: "none",
    onComplete: function() {
    ["Wx"].forEach(tag => {
    const el = document.getElementById(
    `operations-tag-${
    encodeURIComponent(tag)
    }`
    );
    if (el && !el.classList.contains('is-open') && true) { el?.click(); }
    ///**
    });
    }
  });

  //</editor-fold>
};
