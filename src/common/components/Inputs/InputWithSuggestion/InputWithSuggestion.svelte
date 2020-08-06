<script lang="ts">
    import ClickOutside from "svelte-click-outside";

    import {
        Font312DarkGray,
        Font312Black,
        Input,
        FlexVert,
        Wrapper,
        Font312RedAttention,
        InvalidInput,
        Dropdown,
        SingleLine,
        FlexHorCenter,
        DropdownInputExpanded,
        InputWithSuggestion
    } from "./style";

    const onClickOutside = () => {
        isOpen = false;
    }

    const findTopSuggestions = () => {
        if (!value) {
            currentSuggestions = [];
        }

        currentSuggestions = [...suggestionsList.filter(item => item.indexOf(value) != -1)];
        console.log(currentSuggestions)
    }

    const onInputHandler = () => {
        console.log(value)
        if (currentSuggestions) {
            isOpen = true;
        }

        if (!value) {
            return;
        }

        findTopSuggestions();
    }

    //TODO:
    //1. implement real suggestions list
    //2. implement arrow buttons logic
    //3. dry - implement dropdown component
    //4. check chrome
    //5. only 5 top suggestions

    const onSelect = (selectedValue) => {
        value = selectedValue;
        isOpen = false;
    }

    let isOpen = false;
    let currentSuggestions: string[] = [];

    let dropdownElement: HTMLDivElement = null;

    export let value: string = "";
    export let suggestionsList: string[] = [];

    export let autofocus = false;
    export let name = "";
    export let label = "";
    export let disabled = false;
    export let required = false;
    export let invalid = false;
</script>

<div class="{FlexVert} {Wrapper}">
    {#if label}
        <label class="{Font312DarkGray} {invalid && Font312RedAttention}"
               for={name}>{label}</label>
    {/if}

    <ClickOutside on:clickoutside={onClickOutside}>
        <div class={Wrapper}>
            <input autocomplete="off" on:input={onInputHandler} on:change
                   class="{Input} {InputWithSuggestion} {Font312Black} {isOpen && DropdownInputExpanded} {invalid && InvalidInput}"
                   {required}
                   {disabled} {name}
                   bind:value={value} {autofocus}/>
        </div>
        {#if isOpen}
            <div class={Dropdown} bind:this={dropdownElement}>
                {#each currentSuggestions as suggestion}
                    <div on:click={() => onSelect(suggestion)}
                         class="{SingleLine} {FlexHorCenter} {Font312Black}">{suggestion}</div>
                {/each}
            </div>
        {/if}
    </ClickOutside>
</div>

