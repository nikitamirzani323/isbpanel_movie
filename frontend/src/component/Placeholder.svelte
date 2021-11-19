<script>
    export let card_style = "border:none;padding:0px;margin:0px;";
    export let total_placeholder = 1;
    let data = []
    for(let i=0;i<parseInt(total_placeholder);i++){
        data = [
            ...data,
            {
                temp: "asd",
            },
        ];
    }
</script>
<div class="card" aria-hidden="true" style="{card_style}">
    <div class="card-body">
      <h5 class="card-title placeholder-glow">
        <span class="placeholder col-6"></span>
      </h5>
      <p class="card-text placeholder-glow">
          {#each data as rec }
            <span class="placeholder col-7"></span>
            <span class="placeholder col-4"></span>
            <span class="placeholder col-4"></span>
            <span class="placeholder col-6"></span>
            <span class="placeholder col-8"></span>
          {/each }

      </p>
    </div>
</div>