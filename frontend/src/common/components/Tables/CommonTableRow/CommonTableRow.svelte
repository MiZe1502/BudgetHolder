<script lang="typescript">
    import {style} from "./style";

    import {ColumnConfig} from "../CommonTable/utils";

    export let config: ColumnConfig[] = [];
    export let data: Record<any, any> = {};
    export let inPopup: boolean = false;
</script>

<div class="{!inPopup ? style.SideMainPadding : style.PopupPadding} {style.Row} {style.FlexHorCenter}">
    {#each Object.entries(config) as [columnHeader, columnConfig], index}
        <div style={columnConfig.style}
             class="{columnConfig.overflowed && style.Overflowed} {style.FlexHorCenter} {style.Cell}">
            <svelte:component this={columnConfig.component}
                              data={columnConfig.mapping(data)}/>
        </div>
    {/each}
</div>
