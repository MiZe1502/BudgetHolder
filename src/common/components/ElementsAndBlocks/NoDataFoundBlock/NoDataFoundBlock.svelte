<script lang="typescript">
    import { _ } from 'svelte-i18n'
    import { SideMainPadding, Font316DarkGray } from "./style";
</script>

<div class="{SideMainPadding} {Font316DarkGray}">
    {$_("common.messages.no_data.not_found")}
</div>