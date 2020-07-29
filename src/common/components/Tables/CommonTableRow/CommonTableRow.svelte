<script lang="typescript">
    import {
        Row,
        PopupPadding,
        SideMainPadding,
        FlexHorCenter,
        Cell,
        Overflowed
    } from "./style";

    import {ColumnConfig} from "../CommonTable/utils";

    export let config: ColumnConfig[] = [];
    export let data: Record<any, any> = {};
    export let inPopup: boolean = false;
</script>

<div class="{!inPopup ? SideMainPadding : PopupPadding} {Row} {FlexHorCenter}">
    {#each Object.entries(config) as [columnHeader, columnConfig], index}
        <div style={columnConfig.style}
             class="{columnConfig.overflowed && Overflowed} {FlexHorCenter} {Cell}">
            <svelte:component this={columnConfig.component}
                              data={columnConfig.mapping(data)}/>
        </div>
    {/each}
</div>
