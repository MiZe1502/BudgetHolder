<script lang="ts">
    import ClickOutside from "svelte-click-outside";

    import {
        SuggestionItem,
        maxSuggestionsListSize,
        KeyboardKeys
    } from "./utils";
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
        InputWithSuggestion,
        SelectedLine
    } from "./style";

    const onClickOutside = () => {
        isOpen = false;
    }

    const findTopSuggestions = () => {
        if (!value || value.length === 0 || value === "") {
            console.log("here")
            currentSuggestions = [];
        }

        currentSuggestions = [...suggestionsList.filter(item =>
                item.value.trim().toLowerCase().indexOf(value.trim().toLowerCase()) != -1)].slice(0, maxSuggestionsListSize);

        console.log(value, currentSuggestions)
    }

    const onInputHandler = (event) => {
        value = event.target.value;

        findTopSuggestions();

        if (value.length > 0) {
            if (currentSuggestions.length > 0) {
                isOpen = true;
            }
        } else {
            isOpen = false;
        }
    }

    const onKeyHandler = (event: KeyboardEvent<HTMLInputElement>) => {
        if (event.which === KeyboardKeys.ArrowUp) {
            if (currentSelectedSuggestionInList === 0) {
                currentSelectedSuggestionInList = currentSuggestions.length - 1;
                return;
            }
            currentSelectedSuggestionInList -= 1;
            return;
        }
        if (event.which === KeyboardKeys.ArrowDown) {
            if (currentSelectedSuggestionInList === currentSuggestions.length - 1) {
                currentSelectedSuggestionInList = 0
                return;
            }
            currentSelectedSuggestionInList += 1;
            return;
        }
        if (event.which === KeyboardKeys.Enter) {
            onSelect(currentSuggestions[currentSelectedSuggestionInList]);
        }
    }

    //TODO:
    //3. dry - implement dropdown component
    //6. load all goods data to form

    const onSelect = (selectedValue) => {
        value = selectedValue.value;
        onSelectHandler(selectedValue);
        isOpen = false;
    }

    let isOpen = false;
    let currentSuggestions: SuggestionItem[] = [];

    let dropdownElement: HTMLDivElement = null;

    let currentSelectedSuggestionInList: number = 0;

    export let value: string = "";
    export let suggestionsList: SuggestionItem[] = [];

    export let onSelectHandler: (selectedItem: SuggestionItem) => void;
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
            <input autocomplete="off" on:keyup={onKeyHandler}
                   on:input={onInputHandler} on:change
                   class="{Input} {InputWithSuggestion} {Font312Black} {isOpen && DropdownInputExpanded} {invalid && InvalidInput}"
                   {required}
                   {disabled} {name}
                   bind:value={value} {autofocus}/>
        </div>
        {#if isOpen}
            <div class={Dropdown} bind:this={dropdownElement}>
                {#each currentSuggestions as suggestion}
                    <div on:click={() => onSelect(suggestion)}
                         class="{suggestion.id === currentSuggestions[currentSelectedSuggestionInList].id && SelectedLine} {SingleLine} {FlexHorCenter} {Font312Black}">{suggestion.value}</div>
                {/each}
            </div>
        {/if}
    </ClickOutside>
</div>

