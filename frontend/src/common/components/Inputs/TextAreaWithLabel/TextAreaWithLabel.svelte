<script lang="ts">
    import {style} from "./style";

    export let value = "";
    export let autofocus = false;
    export let name = "";
    export let label = "";
    export let disabled = false;
    export let invalid = false;
    export let textAreaClass = "";
</script>

{#if label}
    <label class="{style.Font312DarkGray} {invalid && style.Font312RedAttention}"
           for={name}>{label}</label>
{/if}
<textarea on:input on:change
          class="{style.TextArea} {style.Font312Black} {textAreaClass} {invalid && style.InvalidInput}"
          {disabled} {name}
          bind:value={value}></textarea>
