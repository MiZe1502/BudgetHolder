<script lang="ts">
    import {style} from "./style";

    export let width: number;
    export let height: number;
    export let className: string = "";
</script>

<div class="{style.Dropdown} {className}"
     style={`height: ${height}px; width: ${width}px`}>
    <slot/>
</div>