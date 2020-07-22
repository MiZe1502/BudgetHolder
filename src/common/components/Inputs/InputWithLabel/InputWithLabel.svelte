<script lang="ts">
    import {
        Font312DarkGray,
        Font312Black,
        Input,
        FlexVert,
        Wrapper,
        Font312RedAttention,
        InvalidInput
    } from "./style";

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

<div class="{FlexVert} {Wrapper}">
    {#if label}
        <label class="{Font312DarkGray} {invalid && Font312RedAttention}"
               for={name}>{label}</label>
    {/if}
    <input on:input on:change bind:this={inputElement} on:focus={onFocus}
           class="{Input} {Font312Black} {invalid && InvalidInput}" {required}
           {disabled} {name} {readonly}
           bind:value={value} {autofocus}/>
</div>

