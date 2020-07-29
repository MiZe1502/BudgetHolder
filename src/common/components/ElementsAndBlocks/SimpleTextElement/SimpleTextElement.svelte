<script lang="ts">
    import { Font312Black } from "./style";

    export let data: string = "";
    export let style = "";
</script>

<div class="{Font312Black}" style={style}>{data || "\u2014"}</div>