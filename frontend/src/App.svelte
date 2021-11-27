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
		<!-- Histats.com  START (html only)-->
		<a href="/" alt="page hit counter" target="_blank" >
			<embed src="//s10.histats.com/429.swf"  flashvars="jver=1&acsid=4606934&domi=4"  quality="high"  width="112" height="75" name="429.swf"  align="middle" type="application/x-shockwave-flash" pluginspage="//www.macromedia.com/go/getflashplayer" wmode="transparent" />
		</a>
		<img  src="//sstatic1.histats.com/0.gif?4606934&101" alt="" border="0">
		<!-- Histats.com  END  -->
		<br>
	</center>
{/if}