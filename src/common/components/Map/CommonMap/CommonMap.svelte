<script lang="typescript">
    import {_} from 'svelte-i18n'
    import {style} from "./style";
    import {LoadingStatus} from "../../../../stores/utils";

    import {MapItemData} from "../../ActionElements/MapActionElement/utils";

    import LoadingSpinner
        from '../../ElementsAndBlocks/LoadingSpinner/LoadingSpinner.svelte'
    import NoDataFoundBlock
        from '../../ElementsAndBlocks/NoDataFoundBlock/NoDataFoundBlock.svelte'
    import CommonTitle from "../../Tables/CommonTitle/CommonTitle.svelte";
    import MapContainer from "../MapContainer/MapContainer.svelte";

    export let data: MapItemData[] = [];
    export let status: LoadingStatus = LoadingStatus.None;
</script>

<section class="{style.FlexVert} {style.SectionBottomMargin}">
    <CommonTitle withButton={false} title={$_("common.components.map.title")}
                 dataLength={null}/>
    {#if status === LoadingStatus.Loading}
        <LoadingSpinner/>
    {:else if status === LoadingStatus.Finished}
        {#if !data || data.length === 0}
            <NoDataFoundBlock/>
        {:else}
            <div class="{style.SideMainPadding}">
                <MapContainer data={data}
                              wrapperClass={style.MapContainerWrapper}/>
            </div>
        {/if}
    {:else if status === LoadingStatus.Error}
        {$_("common.messages.errors.fetching")}
    {/if}
</section>