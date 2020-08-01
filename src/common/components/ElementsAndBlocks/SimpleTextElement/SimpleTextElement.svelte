<script lang="ts">
    import { Font312Black } from "./style";

    export let data: string = "";
    export let style = "";
    export let className = "";
</script>

<div class="{Font312Black} {className}" style={style}>{data || "\u2014"}</div>