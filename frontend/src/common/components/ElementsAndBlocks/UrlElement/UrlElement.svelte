<script lang="typescript">
    import {UrlElementData} from "./utils";
    import {style as css} from "./style";

    export let data: UrlElementData = {};
    export let style = "";
</script>

<style>

</style>

<div style={style}
     class="{css.UrlWrapper} {css.Overflowed} {css.FlexHorCenter} {css.Font312Black}">
    {#if data.url}
        <a title={data.name} class="{css.Url} {css.Overflowed}"
           href={data.url}>{data.name}</a>
    {:else}
        <span title={data.name} class="{css.Overflowed}">{data.name}</span>
    {/if}
</div>