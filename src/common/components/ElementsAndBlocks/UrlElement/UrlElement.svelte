<script lang="typescript">
    import {UrlElementData} from "./utils";
    import {style} from "./style";

    export let data: UrlElementData = {};
    export let style = "";
</script>

<style>

</style>

<div style={style}
     class="{style.UrlWrapper} {style.Overflowed} {style.FlexHorCenter} {style.Font312Black}">
    {#if data.url}
        <a title={data.name} class="{style.Url} {style.Overflowed}"
           href={data.url}>{data.name}</a>
    {:else}
        <span title={data.name} class="{style.Overflowed}">{data.name}</span>
    {/if}
</div>