<script lang="typescript">
    import { Row, SideMainPadding, FlexHorCenter } from "./style";

    import { ColumnConfig } from "../CommonTable/utils";

    export let config: ColumnConfig[] = [];
    export let data: Record<any, any> = {};
</script>

<div class="{SideMainPadding} {Row} {FlexHorCenter}">
    {#each Object.entries(config) as [columnHeader, columnConfig], index}
        <svelte:component style={columnConfig.style} this={columnConfig.component} data={columnConfig.mapping(data)}/>
    {/each}
</div>
