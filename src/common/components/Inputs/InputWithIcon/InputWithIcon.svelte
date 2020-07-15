<script lang="ts">
    import { Wrapper, ButtonWrapper, FlexHorCenter } from "./style";

    import InputWithLabel from "../InputWithLabel/InputWithLabel.svelte";

    export let value = "";
    export let autofocus = false;
    export let name = "";
    export let type = "text";
    export let label = "";
    export let disabled = false;
    export let required = false;
    export let onChange = () => {
    };
    export let onInput = () => {
    };
</script>

<div class="{Wrapper} {FlexHorCenter}">
    <InputWithLabel on:input on:change bind:value={value} {autofocus} {name} {type}
                    {label} {disabled} {required}/>
    <div class="{ButtonWrapper}">
        <slot />
    </div>
</div>

