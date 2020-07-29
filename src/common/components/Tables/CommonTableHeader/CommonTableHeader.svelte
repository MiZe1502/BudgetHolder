<script lang="typescript">
    import {_} from 'svelte-i18n'
    import {
        Font716Black,
        Header,
        SideMainPadding,
        FlexHorCenter,
        PopupPadding
    } from "./style";

    import {ColumnConfig} from "../CommonTable/utils";

    export let config: ColumnConfig[] = [];
    export let inPopup: boolean = false;
</script>

<div class="{Font716Black} {!inPopup ? SideMainPadding : PopupPadding} {Header} {FlexHorCenter}">
    {#each config as configItem}
        <span style={configItem.style}>{$_(configItem.header)}</span>
    {/each}
</div>