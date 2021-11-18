<script>
	import Navbar from "./component/Navbar.svelte";
	import Footer from "./component/Footer.svelte";
	import Modal from "./component/Modal.svelte";

	let myModal = "";
	const handleMovie = () => {
		myModal = new bootstrap.Modal(document.getElementById("modalmovie"));
		myModal.show();
	};
	
</script>

<Navbar />
<div class="container" style="margin-top: 50px;">
	<div class="row">
		<div class="col-sm-12" style="margin:0px;padding:0px 3px 0px 0px;">
			<div class="card" style="background-color:#2c2c2c;border:none;margin-top:5px;">
				<div class="card-header" style="padding: 10px 0px 5px 10px;margin:0px;background-color:#2c2c2c;border-bottom:2px solid #ed247a;">
				  <h1 style="font-size: 16px;color:white;font-weight:bold;">Popular</h1>
				</div>
				<div class="card-body" style="margin: 5px;padding:5px;background-color: #2c2c2c;border-bottom:1px solid #2c2c2c;">
					<div class="row">
						<div class="col-sm-2" style="margin:0px;padding:5px;">
							<div
								on:click={() => {
									handleMovie();
								}}  
								class="card" style="background-color:#2c2c2c;border:none;cursor:pointer;">
								<span style="color: white;margin-top: 5px;position:absolute;background-color:rgba(237, 36, 122, 0.3);padding:5px;width:150px;border-top-right-radius:5px;border-bottom-right-radius:5px;">
									Lucifer
								</span>
								<img width="200" src="https://www.themoviedb.org/t/p/w220_and_h330_face/4EYPN5mVIhKLfxGruy7Dy41dTVn.jpg" class="card-img-top" alt="...">
							</div>
						</div>
						<div class="col-sm-2" style="margin:0px;padding:5px;">
							<div
								on:click={() => {
									handleMovie();
							  	}}  
								class="card" style="background-color:#2c2c2c;border:none;cursor:pointer;">
								<span style="color: white;margin-top: 5px;position:absolute;background-color:rgba(237, 36, 122, 0.3);padding:5px;width:150px;border-top-right-radius:5px;border-bottom-right-radius:5px;">
									Lucifer
								</span>
								<img width="150" src="https://www.themoviedb.org/t/p/w220_and_h330_face/5pVJ9SuuO72IgN6i9kMwQwnhGHG.jpg" class="card-img-top" alt="...">
							</div>
						</div>
						<div class="col-sm-2" style="margin:0px;padding:5px;">
							<div
								on:click={() => {
									handleMovie();
								}}  
								class="card" style="background-color:#2c2c2c;border:none;cursor:pointer;">
								<span style="color: white;margin-top: 5px;position:absolute;background-color:rgba(237, 36, 122, 0.3);padding:5px;width:150px;border-top-right-radius:5px;border-bottom-right-radius:5px;">
									Lucifer
								</span>
								<img width="150" src="https://www.themoviedb.org/t/p/w220_and_h330_face/wdE6ewaKZHr62bLqCn7A2DiGShm.jpg" class="card-img-top" alt="...">
							</div>
						</div>
						<div class="col-sm-2" style="margin:0px;padding:5px;">
							<div
								on:click={() => {
									handleMovie();
								}} 
								class="card" style="background-color:#2c2c2c;border:none;cursor:pointer;">
								<span style="color: white;margin-top: 5px;position:absolute;background-color:rgba(237, 36, 122, 0.3);padding:5px;width:150px;border-top-right-radius:5px;border-bottom-right-radius:5px;">
									Lucifer
								</span>
								<img width="150" src="https://www.themoviedb.org/t/p/w220_and_h330_face/1BIoJGKbXjdFDAqUEiA2VHqkK1Z.jpg" class="card-img-top" alt="...">
							</div>
						</div>
						<div class="col-sm-2" style="margin:0px;padding:5px;">
							<div
								on:click={() => {
									handleMovie();
								}} 
								class="card" style="background-color:#2c2c2c;border:none;cursor:pointer;">
								<span style="color: white;margin-top: 5px;position:absolute;background-color:rgba(237, 36, 122, 0.3);padding:5px;width:150px;border-top-right-radius:5px;border-bottom-right-radius:5px;">
									Lucifer
								</span>
								<img width="150" src="https://www.themoviedb.org/t/p/w220_and_h330_face/5YCg0PFnGxDKb24Lmhxz8p4Q8j1.jpg" class="card-img-top" alt="...">
							</div>
						</div>
						<div class="col-sm-2" style="margin:0px;padding:5px;">
							<div
								on:click={() => {
									handleMovie();
								}} 
								class="card" style="background-color:#2c2c2c;border:none;cursor:pointer;">
								<span style="color: white;margin-top: 5px;position:absolute;background-color:rgba(237, 36, 122, 0.3);padding:5px;width:150px;border-top-right-radius:5px;border-bottom-right-radius:5px;">
									Lucifer
								</span>
								<img width="150" src="https://www.themoviedb.org/t/p/w220_and_h330_face/bGZn5RVzMMXju4ev7xbl1aLdXqq.jpg" class="card-img-top" alt="...">
							</div>
						</div>
					</div>
				</div>
			</div>
			<div class="card" style="background-color:#2c2c2c;border:none;margin-top:5px;">
				<div class="card-header" style="padding: 10px 0px 5px 10px;margin:0px;background-color:#2c2c2c;border-bottom:2px solid #ed247a;">
				  <h1 style="font-size: 16px;color:white;font-weight:bold;">Action</h1>
				</div>
				<div class="card-body" style="margin: 5px;padding:5px;background-color: #2c2c2c;border-bottom:1px solid #2c2c2c;">
					<div class="row">
						<div class="col-sm-2" style="margin:0px;padding:5px;">
							<div
								on:click={() => {
									handleMovie();
								}}  
								class="card" style="background-color:#2c2c2c;border:none;cursor:pointer;">
								<span style="color: white;margin-top: 5px;position:absolute;background-color:rgba(237, 36, 122, 0.3);padding:5px;width:150px;border-top-right-radius:5px;border-bottom-right-radius:5px;">
									Lucifer
								</span>
								<img width="200" src="https://www.themoviedb.org/t/p/w220_and_h330_face/mGPdWEEmfzP7VQBQsXrFt1b1ikQ.jpg" class="card-img-top" alt="...">
							</div>
						</div>
						<div class="col-sm-2" style="margin:0px;padding:5px;">
							<div
								on:click={() => {
									handleMovie();
							  	}}  
								class="card" style="background-color:#2c2c2c;border:none;cursor:pointer;">
								<span style="color: white;margin-top: 5px;position:absolute;background-color:rgba(237, 36, 122, 0.3);padding:5px;width:150px;border-top-right-radius:5px;border-bottom-right-radius:5px;">
									Lucifer
								</span>
								<img width="150" src="https://www.themoviedb.org/t/p/w220_and_h330_face/gII53HAH7UA1yx189vROMzWA5ib.jpg" class="card-img-top" alt="...">
							</div>
						</div>
						<div class="col-sm-2" style="margin:0px;padding:5px;">
							<div
								on:click={() => {
									handleMovie();
								}}  
								class="card" style="background-color:#2c2c2c;border:none;cursor:pointer;">
								<span style="color: white;margin-top: 5px;position:absolute;background-color:rgba(237, 36, 122, 0.3);padding:5px;width:150px;border-top-right-radius:5px;border-bottom-right-radius:5px;">
									Lucifer
								</span>
								<img width="150" src="https://www.themoviedb.org/t/p/w220_and_h330_face/clnyhPqj1SNgpAdeSS6a6fwE6Bo.jpg" class="card-img-top" alt="...">
							</div>
						</div>
						<div class="col-sm-2" style="margin:0px;padding:5px;">
							<div
								on:click={() => {
									handleMovie();
								}} 
								class="card" style="background-color:#2c2c2c;border:none;cursor:pointer;">
								<span style="color: white;margin-top: 5px;position:absolute;background-color:rgba(237, 36, 122, 0.3);padding:5px;width:150px;border-top-right-radius:5px;border-bottom-right-radius:5px;">
									Lucifer
								</span>
								<img width="150" src="https://www.themoviedb.org/t/p/w220_and_h330_face/jRcjCpvpZbhtcztNT7xzwq7XcfD.jpg" class="card-img-top" alt="...">
							</div>
						</div>
						<div class="col-sm-2" style="margin:0px;padding:5px;">
							<div
								on:click={() => {
									handleMovie();
								}} 
								class="card" style="background-color:#2c2c2c;border:none;cursor:pointer;">
								<span style="color: white;margin-top: 5px;position:absolute;background-color:rgba(237, 36, 122, 0.3);padding:5px;width:150px;border-top-right-radius:5px;border-bottom-right-radius:5px;">
									Lucifer
								</span>
								<img width="150" src="https://www.themoviedb.org/t/p/w220_and_h330_face/fSmItIOMMURBLBP8ji2nzkhBo3t.jpg" class="card-img-top" alt="...">
							</div>
						</div>
						<div class="col-sm-2" style="margin:0px;padding:5px;">
							<div
								on:click={() => {
									handleMovie();
								}} 
								class="card" style="background-color:#2c2c2c;border:none;cursor:pointer;">
								<span style="color: white;margin-top: 5px;position:absolute;background-color:rgba(237, 36, 122, 0.3);padding:5px;width:150px;border-top-right-radius:5px;border-bottom-right-radius:5px;">
									Lucifer
								</span>
								<img width="150" src="https://www.themoviedb.org/t/p/w220_and_h330_face/4VR2gDn7dYOEfJPfDpVRPEHgqeo.jpg" class="card-img-top" alt="...">
							</div>
						</div>
					</div>
				</div>
			</div>
			<div class="card" style="background-color:#2c2c2c;border:none;margin-top:5px;">
				<div class="card-header" style="padding: 10px 0px 5px 10px;margin:0px;background-color:#2c2c2c;border-bottom:2px solid #ed247a;">
				  <h1 style="font-size: 16px;color:white;font-weight:bold;">Horror</h1>
				</div>
				<div class="card-body" style="margin: 5px;padding:5px;background-color: #2c2c2c;border-bottom:1px solid #2c2c2c;">
					<div class="row">
						<div class="col-sm-2" style="margin:0px;padding:5px;">
							<div
								on:click={() => {
									handleMovie();
								}}  
								class="card" style="background-color:#2c2c2c;border:none;cursor:pointer;">
								<span style="color: white;margin-top: 5px;position:absolute;background-color:rgba(237,36, 122, 0.3);padding:5px;width:150px;border-top-right-radius:5px;border-bottom-right-radius:5px;">
									Lucifer
								</span>
								<img width="200" src="https://www.themoviedb.org/t/p/w220_and_h330_face/a6eIDQKz2uIu69ABt87Evwn1kBP.jpg" class="card-img-top" alt="...">
							</div>
						</div>
						<div class="col-sm-2" style="margin:0px;padding:5px;">
							<div
								on:click={() => {
									handleMovie();
							  	}}  
								class="card" style="background-color:#2c2c2c;border:none;cursor:pointer;">
								<span style="color: white;margin-top: 5px;position:absolute;background-color:rgba(237, 36, 122, 0.3);padding:5px;width:150px;border-top-right-radius:5px;border-bottom-right-radius:5px;">
									Lucifer
								</span>
								<img width="150" src="https://www.themoviedb.org/t/p/w220_and_h330_face/mckCh2GYSerIqyCyPSLIStLvmZF.jpg" class="card-img-top" alt="...">
							</div>
						</div>
						<div class="col-sm-2" style="margin:0px;padding:5px;">
							<div
								on:click={() => {
									handleMovie();
								}}  
								class="card" style="background-color:#2c2c2c;border:none;cursor:pointer;">
								<span style="color: white;margin-top: 5px;position:absolute;background-color:rgba(237, 36, 122, 0.3);padding:5px;width:150px;border-top-right-radius:5px;border-bottom-right-radius:5px;">
									Lucifer
								</span>
								<img width="150" src="https://www.themoviedb.org/t/p/w220_and_h330_face/mWHCII5OWHx5pRSN2VYYLvT8DbB.jpg" class="card-img-top" alt="...">
							</div>
						</div>
						<div class="col-sm-2" style="margin:0px;padding:5px;">
							<div
								on:click={() => {
									handleMovie();
								}} 
								class="card" style="background-color:#2c2c2c;border:none;cursor:pointer;">
								<span style="color: white;margin-top: 5px;position:absolute;background-color:rgba(237, 36, 122, 0.3);padding:5px;width:150px;border-top-right-radius:5px;border-bottom-right-radius:5px;">
									Lucifer
								</span>
								<img width="150" src="https://www.themoviedb.org/t/p/w220_and_h330_face/ykcL3gNhckBgm54f1y0r97k8dJ.jpg" class="card-img-top" alt="...">
							</div>
						</div>
						<div class="col-sm-2" style="margin:0px;padding:5px;">
							<div
								on:click={() => {
									handleMovie();
								}} 
								class="card" style="background-color:#2c2c2c;border:none;cursor:pointer;">
								<span style="color: white;margin-top: 5px;position:absolute;background-color:rgba(237, 36, 122, 0.3);padding:5px;width:150px;border-top-right-radius:5px;border-bottom-right-radius:5px;">
									Lucifer
								</span>
								<img width="150" src="https://www.themoviedb.org/t/p/w220_and_h330_face/qZMEiTsNlCQV27hHQE27ZtlPWyv.jpg" class="card-img-top" alt="...">
							</div>
						</div>
						<div class="col-sm-2" style="margin:0px;padding:5px;">
							<div
								on:click={() => {
									handleMovie();
								}} 
								class="card" style="background-color:#2c2c2c;border:none;cursor:pointer;">
								<span style="color: white;margin-top: 5px;position:absolute;background-color:rgba(237, 36, 122, 0.3);padding:5px;width:150px;border-top-right-radius:5px;border-bottom-right-radius:5px;">
									Lucifer
								</span>
								<img width="150" src="https://www.themoviedb.org/t/p/w220_and_h330_face/mI6mFCO3lrJ4ovzwfmoLs58uhlK.jpg" class="card-img-top" alt="...">
							</div>
						</div>
					</div>
				</div>
			</div>
		</div>
	</div>
</div>
<Footer />
<Modal
  modal_id="modalmovie"
  modal_size="modal-dialog-centered modal-xl"
  modal_title="Movie"
  modal_modal_css="background-color:#191c1f;border:none;"
  modal_header_css="color:white;font-weight:bold;"
  modal_headertitle_css="font-size:14px;"
  modal_body_css="padding:5px;"
  modal_footer_css="padding:5px;"
  modal_footer={false}
  header_flag={false}>
  <slot:template slot="header">
    
  </slot:template>
  <slot:template slot="body">
    <div class="ratio ratio-16x9">
		<iframe 
			src="https://dood.ws/e/pkfcf4ir009p" 
			allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture" 
			title="YouTube video" allowfullscreen></iframe>
	  </div>
  </slot:template>
</Modal>