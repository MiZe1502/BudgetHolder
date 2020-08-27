<script lang="ts">
    import {style} from "./style";

    export let value = "";
    export let autofocus = false;
    export let name = "";
    export let type = "text";
    export let label = "";
    export let disabled = false;
    export let required = false;
    export let invalid = false;
    export let readonly = false;

    let inputElement: HTMLInputElement = null;

    //dirty hack to set type from props
    const onFocus = () => {
        inputElement.setAttribute("type", type)
    }
</script>

<div class="{style.FlexVert} {style.Wrapper}">
    {#if label}
        <label class="{style.Font312DarkGray} {invalid && style.Font312RedAttention}"
               for={name}>{label}</label>
    {/if}
    <input on:input on:change bind:this={inputElement} on:focus={onFocus}
           class="{style.Input} {style.Font312Black} {invalid && style.InvalidInput}"
           {required}
           {disabled} {name} {readonly}
           bind:value={value} {autofocus}/>
</div>

