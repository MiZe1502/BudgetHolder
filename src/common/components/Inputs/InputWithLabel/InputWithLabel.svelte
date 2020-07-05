<script lang="ts">
    import {Font312DarkGray, Font316Black, Input} from "./style";

    export let value = "";
    export let autofocus = false;
    export let name = "";
    export let type = "text";
    export let label = "";
    export let disabled = false;

    let inputElement: HTMLInputElement = null;

    //dirty hack to set type from props
    const onFocus = () => {
        inputElement.setAttribute("type", type)
    }
</script>

{#if label}
    <label class="{Font312DarkGray}" for={name}>{label}</label>
{/if}
<input bind:this={inputElement} on:focus={onFocus}
       class="{Input} {Font316Black}" {disabled} {name} bind:value={value} {autofocus}/>
