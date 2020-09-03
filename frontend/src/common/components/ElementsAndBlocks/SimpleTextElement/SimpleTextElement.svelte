<script lang="ts">
    import {style as css} from "./style";

    export let data: string = "";
    export let style = "";
    export let className = "";
</script>

<div class="{css.Font312Black} {className}"
     style={style}>{data || "\u2014"}</div>