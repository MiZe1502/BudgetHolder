<script lang="typescript">
    import {_} from 'svelte-i18n'
    import {style} from "./style";

    import {ColumnConfig} from "../CommonTable/utils";

    export let config: ColumnConfig[] = [];
    export let inPopup: boolean = false;
</script>

<div class="{style.Font716Black} {!inPopup ? style.SideMainPadding : style.PopupPadding} {style.Header} {style.FlexHorCenter}">
    {#each config as configItem}
        <span style={configItem.style}>{$_(configItem.header)}</span>
    {/each}
</div>