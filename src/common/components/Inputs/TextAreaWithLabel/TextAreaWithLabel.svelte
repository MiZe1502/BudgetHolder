<script lang="ts">
    import {Font312DarkGray, Font312Black, TextArea} from "./style";

    export let value = "";
    export let autofocus = false;
    export let name = "";
    export let label = "";
    export let disabled = false;
    export let textAreaClass = "";
    export let onChange = () => {};
    export let onInput = () => {};
</script>

{#if label}
    <label class="{Font312DarkGray}" for={name}>{label}</label>
{/if}
<textarea on:input={onInput} on:change={onChange} class="{TextArea} {Font312Black} {textAreaClass}" {disabled} {name}
          bind:value={value}></textarea>
