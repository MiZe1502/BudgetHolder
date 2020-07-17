<script lang="typescript">
    import { SideMainPadding, FlexVert, SectionBottomMargin, MapContainerWrapper } from "./style";
    import { LoadingStatus } from "../../../../stores/utils";

    import { MapItemData } from "../../ActionElements/MapActionElement/utils";

    import LoadingSpinner from '../../ElementsAndBlocks/LoadingSpinner/LoadingSpinner.svelte'
    import NoDataFoundBlock from '../../ElementsAndBlocks/NoDataFoundBlock/NoDataFoundBlock.svelte'
    import CommonTitle from "../../Tables/CommonTitle/CommonTitle.svelte";
    import MapContainer from "../MapContainer/MapContainer.svelte";

    export let data: MapItemData[] = [];
    export let status: LoadingStatus = LoadingStatus.None;
</script>

<section class="{FlexVert} {SectionBottomMargin}">
    <CommonTitle withButton={false} title={"Map"} dataLength={null} />
    {#if status === LoadingStatus.Loading}
        <LoadingSpinner />
    {:else if status === LoadingStatus.Finished}
        {#if !data || data.length === 0}
            <NoDataFoundBlock />
        {:else}
            <div class="{SideMainPadding}">
                <MapContainer data={data} wrapperClass={MapContainerWrapper}/>
            </div>
        {/if}
    {:else if status === LoadingStatus.Error}
        Error fetching data
    {/if}
</section>