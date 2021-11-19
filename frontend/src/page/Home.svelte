<script>
    import Modal from "../component/Modalcustom.svelte";
    import Placeholder from "../component/Placeholder.svelte";
    
    let listmovie = [];
    let movie_title = "";
    let movie_src = "";
    let myModal = "";
	const loaded = new Map();
    async function call_movie() {
		const resdata = await fetch("/api/listmovie", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
            },
            body: JSON.stringify({}),
            });
            if (!resdata.ok) {
            const pasarMessage = `An error has occured: ${resdata.status}`;
            throw new Error(pasarMessage);
            } else {
            const jsondata = await resdata.json();
            if (jsondata.status == 200) {
                let record = jsondata.record;
                if (record != null) {
                for (var i = 0; i < record.length; i++) {
                    listmovie = [
                    ...listmovie,
                        {
                        
                            movie_idcategory: record[i]["movie_idcategory"],
                            movie_category: record[i]["movie_category"],
                            movie_list: record[i]["movie_list"],
                        },
                    ];
                }
                } else {
                alert("Error");
                }
            } 
		}
	}
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
			}, 1000);
		}

		return {
			destroy(){} // noop
		};
	}
    const handleMovie = (src,title) => {
        movie_title = title
        movie_src = src
		myModal = new bootstrap.Modal(document.getElementById("modalmovie"));
		myModal.show();
	};
    const handleCloseModal = () => {
        movie_title = ""
        movie_src = ""
		myModal.hide();
	};
    call_movie()
</script>
<div class="col-sm-12" style="margin:0px;padding:0px 3px 0px 0px;">
    {#if listmovie != ""}
    {#each listmovie as rec}
        <div class="card" style="background-color:#2c2c2c;border:none;margin-top:5px;">
            <div class="card-header" style="padding: 10px 0px 5px 10px;margin:0px;background-color:#2c2c2c;border-bottom:2px solid #ed247a;">
            <h1 style="font-size: 16px;color:white;font-weight:bold;">{rec.movie_category}</h1>
            </div>
            <div class="card-body" style="margin: 5px;padding:5px;background-color: #2c2c2c;border-bottom:1px solid #2c2c2c;">
                <div class="row">
                    {#each rec.movie_list as rec2}
                        <div class="col-sm-2" style="margin:0px;padding:5px;">
                            <div
                                on:click={() => {
                                    handleMovie(rec2.movie_video,rec2.movie_title);
                                }}  
                                class="card" style="background-color:#2c2c2c;border:none;cursor:pointer;">
                                <span style="color: white;margin-top: 5px;position:absolute;background-color:rgba(237, 36, 122, 0.3);padding:5px;width:150px;border-top-right-radius:5px;border-bottom-right-radius:5px;">
                                    {rec2.movie_label}
                                </span>
                                <img
                                    width="200"
                                    alt="{rec2.movie_title}"
                                    src="placeholder.png"
                                    use:lazy="{{src: rec2.movie_thumbnail}}">
                            </div>
                        </div>
                    {/each}
                </div>
            </div>
        </div>
    {/each}
    {:else}
        <Placeholder total_placeholder="3" card_style="background-color:#2c2c2c;border:none;margin-top:5px;" />
    {/if}
</div>
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
  </slot:template>
</Modal>