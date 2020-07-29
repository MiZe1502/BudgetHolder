<script lang="ts">
    import {Font312Black} from "./style";
    import {dateToDMY} from "./utils";

    export let data: Date = null
    export let style = ""
</script>

<div class="{Font312Black}" style={style}>{dateToDMY(data) || "\u2014"}</div>