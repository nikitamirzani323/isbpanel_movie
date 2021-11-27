<script>
	import Router from "svelte-spa-router";
	import { wrap } from "svelte-spa-router/wrap";
	import Navbar from "./component/Navbar.svelte";
	import Footer from "./component/Footer.svelte";
	import Home from "./page/Home.svelte";
	
	let client_device = "";
	let client_token = localStorage.getItem("token");
	console.log(client_token)
	let listmovie = [];
	if(/Android|webOS|iPhone|iPad|iPod|BlackBerry|IEMobile|Opera Mini/i.test(navigator.userAgent)) {
    	client_device = "MOBILE";
	}else{
    	client_device = "WEBSITE";
	}
	async function initapp() {
		const resInit = await fetch("/api/initmovie", {
			method: "POST",
			headers: {
				"Content-Type": "application/json",
			},
			body: JSON.stringify({
			}),
		});
		if (!resInit.ok) {
			const initMessage = `An error has occured: ${resInit.status}`;
			throw new Error(initMessage);
		} else {
			const initJson = await resInit.json();
			if (initJson.status === 200) {
				localStorage.setItem("token", initJson.token);
				call_movie(initJson.token)
			}
		}
	}
	async function call_movie(e) {
		listmovie = [];
		const resdata = await fetch("/api/listmovie", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
                Authorization: "Bearer " + e,
            },
            body: JSON.stringify({}),
            });
            if (!resdata.ok) {
                const pasarMessage = `An error has occured: ${resdata.status}`;
                throw new Error(pasarMessage);
            }else{
                const jsondata = await resdata.json();
                if (jsondata.status == 200) {
                    let record = jsondata.record;
                    if (record != null) {
                        for (var i = 0; i < record.length; i++) {
                            listmovie = [
                            ...listmovie,
                                {
                                    movie_idcategory: record[i]["movie_idcategory"],
                                    movie_type: record[i]["movie_type"],
                                    movie_category: record[i]["movie_category"],
                                    movie_list: record[i]["movie_list"],
                                },
                            ];
                        }
                    } else {
                        alert("Error");
                    }
                }else if(jsondata.status == 400){
                    localStorage.removeItem("token");
					window.location.href = "/";
                } 
		}
	}
	if(client_token == null){
		initapp()
	}else{
		call_movie(client_token)
	}
</script>
{#if client_device =="WEBSITE"}
	{#if client_token !=""}
	<Navbar />
	
		<div class="container" style="margin-top: 50px;">
			<div class="row">
				<Home
					{listmovie}
					{client_token} />
			</div>
		</div>
	
	<Footer />
	{/if}
{:else}
<Navbar />
	<center style="margin-top: 200px;">
		<a style="margin-top: 200px;" href="https://play.google.com/store/apps/details?id=com.isb.tv">
			<img width="150" src="https://duniafilm.b-cdn.net/img/download.png" alt="ISBFILM">
		</a>
		<br>
		 <!-- Histats.com  (div with counter) --><div id="histats_counter"></div>
<!-- Histats.com  START  (aync)-->
<script type="text/javascript">var _Hasync= _Hasync|| [];
    _Hasync.push(['Histats.start', '1,4606934,4,429,112,75,00011111']);
    _Hasync.push(['Histats.fasi', '1']);
    _Hasync.push(['Histats.track_hits', '']);
    (function() {
    var hs = document.createElement('script'); hs.type = 'text/javascript'; hs.async = true;
    hs.src = ('//s10.histats.com/js15_as.js');
    (document.getElementsByTagName('head')[0] || document.getElementsByTagName('body')[0]).appendChild(hs);
    })();</script>
    <noscript><a href="/" target="_blank"><img  src="//sstatic1.histats.com/0.gif?4606934&101" alt="" border="0"></a></noscript>
    <!-- Histats.com  END  -->
		<br>
	</center>
{/if}