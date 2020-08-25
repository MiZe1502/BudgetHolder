<script lang="ts">
    import {Dropdown} from "./style";

    export let width: number;
    export let height: number;
    export let className: string = "";
</script>

<div class="{Dropdown} {className}"
     style={`height: ${height}px; width: ${width}px`}>
    <slot/>
</div>