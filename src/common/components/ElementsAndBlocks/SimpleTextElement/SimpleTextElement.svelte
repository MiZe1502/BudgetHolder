<script lang="ts">
    import {style} from "./style";

    export let data: string = "";
    export let style = "";
    export let className = "";
</script>

<div class="{style.Font312Black} {className}"
     style={style}>{data || "\u2014"}</div>