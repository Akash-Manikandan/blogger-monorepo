<script>
	let { text = '' } = $props();
</script>

<div class="text-overflow" title={text}>
	{text}
</div>

<style>
	.text-overflow {
		overflow: hidden;
		white-space: nowrap;
		text-overflow: ellipsis;
		display: block;
	}
</style>
