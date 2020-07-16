<script lang="ts">
    import {
        Font312DarkGray,
        Font312RedAttention,
        Font312Black,
        TextArea,
        InvalidInput
    } from "./style";

    export let value = "";
    export let autofocus = false;
    export let name = "";
    export let label = "";
    export let disabled = false;
    export let invalid = false;
    export let textAreaClass = "";
</script>

{#if label}
    <label class="{Font312DarkGray} {invalid && Font312RedAttention}"
           for={name}>{label}</label>
{/if}
<textarea on:input on:change
          class="{TextArea} {Font312Black} {textAreaClass} {invalid && InvalidInput}"
          {disabled} {name}
          bind:value={value}></textarea>
