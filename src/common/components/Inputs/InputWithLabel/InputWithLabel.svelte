<script lang="ts">
    import {Font312DarkGray, Font312Black, Input, FlexVert, Wrapper} from "./style";

    export let value = "";
    export let autofocus = false;
    export let name = "";
    export let type = "text";
    export let label = "";
    export let disabled = false;
    export let required = false;
    export let onChange = () => {
    };

    let inputElement: HTMLInputElement = null;

    //dirty hack to set type from props
    const onFocus = () => {
        inputElement.setAttribute("type", type)
    }
</script>

<div class="{FlexVert} {Wrapper}">
    {#if label}
        <label class="{Font312DarkGray}" for={name}>{label}</label>
    {/if}
    <input on:change={onChange} bind:this={inputElement} on:focus={onFocus}
           class="{Input} {Font312Black}" {required} {disabled} {name}
           bind:value={value} {autofocus}/>
</div>

