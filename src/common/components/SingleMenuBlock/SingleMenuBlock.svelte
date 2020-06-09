<script lang="ts">
    import { Link } from "svelte-routing";

    export let title = "";
    export let url = "";
</script>

<style lang="scss">
    .SingleMenuBlock {
        height: 5vh;
        width: auto;
        background-color: #80cf47;
        padding-bottom: 8px;
        padding-top: 8px;
        padding-left: 12px;
        padding-right: 12px;
    }
</style>


<div class="SingleMenuBlock Font312Black">
    <Link to={url}>{title}</Link>
</div>