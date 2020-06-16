<script lang="typescript">
    import { Row, SideMainPadding, FlexHorCenter } from "./style";

    import { ColumnConfig } from "../CommonTable/utils";

    export let config: ColumnConfig[] = [];
    export let data: Record<any, any> = {};
</script>

<div class="{SideMainPadding} {Row} {FlexHorCenter}">
    {#each Object.entries(data) as [fieldName, fieldData], index}
            <svelte:component style={config[index].style} this={config[index].component} data={fieldData}/>
    {/each}
</div>
