<script lang="typescript">
    import { UrlElementData } from "./utils";
    import { Font312Black, FlexHorCenter, Overflowed, UrlWrapper, Url } from "./style";

    export let data: UrlElementData = {};
    export let style = "";
</script>

<style>

</style>

<div style={style} class="{UrlWrapper} {Overflowed} {FlexHorCenter} {Font312Black}">
    {#if data.url}
        <a title={data.name} class="{Url} {Overflowed}" href={data.url}>{data.name}</a>
    {:else}
        <span title={data.name} class="{Overflowed}">{data.name}</span>
    {/if}
</div>