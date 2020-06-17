<script lang="typescript">
    import { UrlElementData } from "./utils";

    export let data: UrlElementData = {};
    export let style = "";
</script>


<style>
    .Wrapper {
        display: flex;
        flex-direction: row;
        align-items: center;
        max-width: 100%;
    }

    .Overflowed {
        overflow: hidden;
        white-space: nowrap;
        text-overflow: ellipsis;
    }
</style>

<div style={style} class="Wrapper Overflowed">
    {#if data.url}
        <a href={data.url}>{data.name}</a>
    {:else}
        <span>{data.name}</span>
    {/if}
</div>