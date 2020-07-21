<script lang="typescript">
    import { _ } from 'svelte-i18n'
    import { Font716Black, Header, SideMainPadding, FlexHorCenter} from "./style";

    import { ColumnConfig } from "../CommonTable/utils";

    export let config: ColumnConfig[] = [];
</script>

<div class="{Font716Black} {SideMainPadding} {Header} {FlexHorCenter}">
    {#each config as configItem}
        <span style={configItem.style}>{$_(configItem.header)}</span>
    {/each}
</div>