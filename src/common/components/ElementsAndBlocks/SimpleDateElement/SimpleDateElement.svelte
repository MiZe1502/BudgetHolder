<script lang="ts">
    import {style as css} from "./style";
    import {dateToDMY} from "./utils";

    export let data: Date = null
    export let style = ""
</script>

<div class="{css.Font312Black}"
     style={style}>{dateToDMY(data) || "\u2014"}</div>