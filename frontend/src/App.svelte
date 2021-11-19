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
{:else}
	<center style="margin-top: 50px;">
		<h1></h1>
		<a href="https://play.google.com/store/apps/details?id=com.isb.tv">
			<img width="150" src="https://duniafilm.b-cdn.net/img/download.png" alt="">
		</a>
	</center>
{/if}