<script lang="typescript">
    import { Font716Black, Header, SideMainPadding, FlexHorCenter} from "./style";

    import { ColumnConfig } from "../CommonTable/utils";

    export let config: ColumnConfig[] = [];
</script>

<div class="{Font716Black} {SideMainPadding} {Header} {FlexHorCenter}">
    {#each config as configItem}
        <span style={configItem.style}>{configItem.header}</span>
    {/each}
</div>