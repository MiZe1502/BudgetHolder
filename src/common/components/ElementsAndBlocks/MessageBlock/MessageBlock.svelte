<script lang="ts">
    import {style} from "./style";

    export let messageText: string;
    export let height: number;
</script>

<div style="height: ${height}" class="{style.SideMainPadding} {style.FlexHorCenter} {style.Message} {style.Font732Black}">
    {messageText}
</div>