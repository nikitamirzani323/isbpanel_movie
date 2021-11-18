<script>
    export let header = true;
    export let footer = true;
    export let card_style = "border:none;padding:0px;margin:0px;";
    export let header_style = "";
    export let body_style = "";
</script>
<div class="card" style="{card_style}">
    {#if header}
        <div class="card-header" style="{header_style}">
            <slot name="header" />
        </div>
    {/if}
    <div id="stream" data-simplebar style={body_style}>
        <slot name="body" />
    </div>
    {#if footer}
    <div class="card-footer">
        <slot name="footer" />
    </div>
    {/if}
</div>


<style>
    #stream::-webkit-scrollbar {
        width: 0.3em;
    }
    #stream::-webkit-scrollbar-track {
        -webkit-box-shadow: inset 0 0 6px rgba(0, 0, 0, 0.3);
    }
    #stream::-webkit-scrollbar-thumb {
        background-color: #505768;
        outline: 1px solid slategrey;
    }
</style>