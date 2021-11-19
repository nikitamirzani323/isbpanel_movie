<script>
	import Router from "svelte-spa-router";
	import { wrap } from "svelte-spa-router/wrap";
	import Navbar from "./component/Navbar.svelte";
	import Footer from "./component/Footer.svelte";
	import Home from "./page/Home.svelte";
	
	let client_device = "";
	if(/Android|webOS|iPhone|iPad|iPod|BlackBerry|IEMobile|Opera Mini/i.test(navigator.userAgent)) {
    	// true for mobile device
    	client_device = "MOBILE";
	}else{
    	client_device = "WEBSITE";
	}
	console.log(client_device)

	
</script>
{#if client_device =="WEBSITE"}
<Navbar />

<div class="container" style="margin-top: 50px;">
	<div class="row">
		<Home />
		
	</div>
</div>
<Footer />
{/if}