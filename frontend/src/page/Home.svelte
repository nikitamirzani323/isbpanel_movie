<script>
    import Modal from "../component/Modalcustom.svelte";
    import Placeholder from "../component/Placeholder.svelte";

    export let client_token = "";
    export let listmovie = [];
    let movie_title = "";
    let movie_srclist= [];
    let movie_src = "";
    let myModal = "";
    let listseason = [];
    let listepisode = [];
    let season_select = "";
	const loaded = new Map();
    
    function lazy(node, data) {
		if (loaded.has(data.src)) {
			node.setAttribute('src', data.src);
		} else {
			// simulate slow loading network
			setTimeout(() => {
				const img = new Image();
				img.src = data.src;
				img.onload = () => {
					loaded.set(data.src, img);
					node.setAttribute('src', data.src); 
				};
			}, 100);
		}

		return {
			destroy(){} // noop
		};
	}
    const handleMovie = (idmovie,src,title,tipe) => {
        movie_title = title
        movie_srclist= [];
        let no = 0
        if(tipe == "movie"){
            if(src != null){
                if(src.length > 1){
                    for(let i=0; i<src.length;i++){
                        no = no + 1
                        movie_src = src[i]["movie_src"]
                        movie_srclist = [
                        ...movie_srclist,
                            {
                                movie_no: no,
                                movie_src: src[i]["movie_src"],
                            },
                        ];
                    }
                }else{
                    movie_src = src[0]["movie_src"]
                }
            }
            myModal = new bootstrap.Modal(document.getElementById("modalmovie"));
		    myModal.show();
        }else{
            call_season(idmovie)
           
            myModal = new bootstrap.Modal(document.getElementById("modalseries"));
		    myModal.show();
        }
        console.log(movie_srclist)
	};
    const handleChangeSource = (e) => {
        movie_src = e
	};
    const handleCloseModal = () => {
        movie_title = ""
        movie_src = ""
        movie_srclist = []
        listseason = []
        listepisode = []
		myModal.hide();
	};
    const handleClickMenu = (e) => {
        // page = e.detail.page
        console.log(e.detail)
    };
    const handleSelectSeason = (event) => {
        call_episode(event.target.value);
    };
    const handleSelectEpisode = (event) => {
        movie_src = event.target.value
    };
    async function call_season(e) {
		listseason = [];
		const resdata = await fetch("/api/listseason", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
                Authorization: "Bearer " + client_token,
            },
            body: JSON.stringify({
                movie_id: parseInt(e)
            }),
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
                        listseason = [
                        ...listseason,
                            {
                                season_id: record[i]["season_id"],
                                season_title: record[i]["season_title"],
                            },
                        ];
                    }
                } else {
                    alert("Error");
                }
            }
		}
	}
    async function call_episode(e) {
		listepisode = [];
		const resdata = await fetch("/api/listepisode", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
                Authorization: "Bearer " + client_token,
            },
            body: JSON.stringify({
                season_id: parseInt(e)
            }),
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
                        if(parseInt(i) == 0){
                            movie_src = record[i]["episode_src"]
                        }
                        listepisode = [
                        ...listepisode,
                            {
                                episode_id: record[i]["episode_id"],
                                episode_title: record[i]["episode_title"],
                                episode_src: record[i]["episode_src"],
                            },
                        ];
                    }
                } else {
                    alert("Error");
                }
            }
		}
	}
</script>
<div class="col-sm-12" style="margin:0px;padding:0px 3px 0px 0px;">
    {#if listmovie != ""}
        {#each listmovie as rec}
            <div class="card" style="background-color:#1e152e;border:none;margin-top:5px;">
                <div class="card-header" style="padding: 10px 0px 5px 10px;margin:0px;background-color:#1e152e;border-bottom:2px solid #d8278f;">
                    <h1 style="font-size: 16px;color:white;font-weight:bold;">{rec.movie_category}</h1>
                </div>
                <div class="card-body" style="margin: 5px;padding:5px;background-color: #1e152e;border-bottom:1px solid #1e152e;">
                    <div class="row">
                        {#each rec.movie_list as rec2}
                            <div class="col-sm-3 col-md-1 col-lg-1 col-xl-1" style="margin:0px;padding:3px;">
                                <div
                                    on:click={() => {
                                        handleMovie(rec2.movie_id,rec2.movie_video,rec2.movie_title,rec2.movie_type);
                                    }}  
                                    class="card" style="background-color:#1e152e;border:none;cursor:pointer;">
                                    
                                    <img
                                        style="border: 1px solid #1e152e;background-color: none;"
                                        class="img-thumbnail"
                                        alt="{rec2.movie_title}"
                                        src="https://imagedelivery.net/W-Usm3AjeE17sxpltvGRNA/fd0287a2-353d-4b47-9a6c-9c8df2ab3f00/public"
                                        use:lazy="{{src: rec2.movie_thumbnail}}">
                                </div>
                            </div>
                        {/each}
                    </div>
                </div>
            </div>
        {/each}
    {:else}
        <Placeholder total_placeholder="6" card_style="background-color:#2c2c2c;border:none;margin-top:5px;" />
        <Placeholder total_placeholder="4" card_style="background-color:#2c2c2c;border:none;margin-top:5px;" />
        <Placeholder total_placeholder="6" card_style="background-color:#2c2c2c;border:none;margin-top:5px;" />
    {/if}
</div>
<style>
    .img-thumbnail {
        padding: 0;
        background-color: #2c2c2c;
        border: 1px solid #2c2c2c;
        border-radius: .2rem;
        max-width: 100%;
        height: auto;
    }
</style>
<Modal 
  modal_id="modalmovie"
  modal_size="modal-dialog-centered modal-xl"
  modal_title="{movie_title}"
  modal_modal_css="background-color:#191c1f;border:none;"
  modal_header_css="color:white;font-weight:bold;"
  modal_headertitle_css="font-size:14px;"
  modal_body_css="padding:5px;"
  modal_footer_css="padding:5px;"
  modal_footer={false}
  header_flag={false}>
  <slot:template slot="modal-header">
    <button
        on:click={() => {
            handleCloseModal();
        }}
        type="button"
        class="btn-close btn-close-white"
        aria-label="Close"></button>
  </slot:template>
  <slot:template slot="body">
    <div class="ratio ratio-16x9">
		<iframe 
			src="{movie_src}" 
			allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture" 
			title="YouTube video" allowfullscreen></iframe>
	</div>
    {#each movie_srclist as rec}
    <button
        on:click={() => {
            handleChangeSource(rec.movie_src);
        }} 
        type="button" class="btn btn-warning btn-sm">SERVER {rec.movie_no}</button>&nbsp;
    {/each}
  </slot:template>
</Modal>
<Modal
    on:handleClickMenu={handleClickMenu}  
    modal_id="modalseries"
    modal_size="modal-dialog-centered modal-xl"
    modal_title="{movie_title}"
    modal_modal_css="background-color:#191c1f;border:none;"
    modal_header_css="color:white;font-weight:bold;"
    modal_headertitle_css="font-size:14px;"
    modal_body_css="padding:5px;"
    modal_footer_css="padding:5px;"
    modal_footer={false}
    header_flag={false}>
    menu={false}>
    <slot:template slot="modal-header">
        <button
            on:click={() => {
                handleCloseModal();
            }}
            type="button"
            class="btn-close btn-close-white"
            aria-label="Close"></button>
    </slot:template>
    <slot:template slot="body">
        <div class="row">
            <div class="col-sm-2">
                <div class="mb-3">
                    <label for="exampleForm" class="form-label">Season</label>
                    <select
                        on:change={handleSelectSeason}
                        bind:value="{season_select}"
                        style="background-color: #222222;color:white;border:none;"
                        class="form-control">
                        {#each listseason as rec }
                        <option value="{rec.season_id}">{rec.season_title}</option>
                        {/each}
                    </select>
                </div>
                <div class="mb-3">
                    <label for="exampleForm" class="form-label">Episode</label>
                    <select
                        on:change={handleSelectEpisode}
                        style="background-color: #222222;color:white;border:none;"
                        class="form-control" 
                        name="" id="">
                        {#each listepisode as rec }
                        <option value="{rec.episode_src}">{rec.episode_title}</option>
                        {/each}
                    </select>
                </div>
            </div>
            <div class="col-sm-10">
                <div class="ratio ratio-16x9">
                    <iframe 
                        src="{movie_src}" 
                        allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture" 
                        title="YouTube video" allowfullscreen></iframe>
                </div>
            </div>
        </div>
    </slot:template>
</Modal>