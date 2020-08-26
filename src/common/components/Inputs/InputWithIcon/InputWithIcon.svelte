<script lang="ts">
    import { style } from "./style";

    import InputWithLabel from "../InputWithLabel/InputWithLabel.svelte";

    export let value = "";
    export let autofocus = false;
    export let name = "";
    export let type = "text";
    export let label = "";
    export let disabled = false;
    export let required = false;
    export let invalid = false;
</script>

<div class="{style.Wrapper} {style.FlexHorCenter}">
    <InputWithLabel on:input on:change bind:value={value} {autofocus} {name} {type}
                    {label} {disabled} {required} {invalid}/>
    <div class="{style.ButtonWrapper}">
        <slot />
    </div>
</div>

