<script>
    import { link } from "svelte-spa-router";
</script>
<div class="container">
    <nav class="navbar fixed-top navbar-light d-flex flex-wrap justify-content-center py-2 mb-5 " style="height:50px;background-color:#1e1e1e;">
        <a href="/" class="d-flex align-items-center mb-2 mb-lg-0 text-white text-decoration-none">
            <img style="margin-top:-5px;" src="logo_movie.png" width="100" alt="" />
      </a>

      <ul class="nav col-12 col-lg-auto me-lg-auto mb-2 justify-content-center mb-md-0">
          <li><a use:link href="/"  class="nav-link px-2 text-secondary">Beranda</a></li>
          <li><a use:link href="/series" class="nav-link px-2 text-white">Acara TV</a></li>
          <li><a use:link href="/movie"  class="nav-link px-2 text-white">Film</a></li>
        </ul>
    </nav>
</div>