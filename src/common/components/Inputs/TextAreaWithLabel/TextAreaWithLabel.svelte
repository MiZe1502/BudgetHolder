<script lang="ts">
    import {Font312DarkGray, Font316Black, TextArea} from "./style";

    export let value = "";
    export let autofocus = false;
    export let name = "";
    export let label = "";
    export let disabled = false;
    export let textAreaClass = "";
</script>

{#if label}
    <label class="{Font312DarkGray}" for={name}>{label}</label>
{/if}
<textarea class="{TextArea} {Font316Black} {textAreaClass}" {disabled} {name}
          bind:value={value}></textarea>
