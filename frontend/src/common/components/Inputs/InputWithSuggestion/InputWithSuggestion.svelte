<script lang="ts">
    import ClickOutside from "svelte-click-outside";

    import {
        SuggestionItem,
        maxSuggestionsListSize,
        KeyboardKeys
    } from "./utils";
    import {style} from "./style";
    import DropdownMenu from "../../DropdownMenu/DropdownMenu.svelte";

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
            isOpen = currentSuggestions.length > 0;
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

    const onSelect = (selectedValue) => {
        value = selectedValue.value;
        onSelectHandler(selectedValue);
        isOpen = false;
    }

    let isOpen = false;
    let currentSuggestions: SuggestionItem[] = [];

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

<div class="{style.FlexVert} {style.Wrapper}">
    {#if label}
        <label class="{style.Font312DarkGray} {invalid && style.Font312RedAttention}"
               for={name}>{label}</label>
    {/if}

    <ClickOutside on:clickoutside={onClickOutside}>
        <div class={style.Wrapper}>
            <input autocomplete="off" on:keyup={onKeyHandler}
                   on:input={onInputHandler} on:change
                   class="{style.Input} {style.InputWithSuggestion} {style.Font312Black} {isOpen && style.DropdownInputExpanded} {invalid && style.InvalidInput}"
                   {required}
                   {disabled} {name}
                   bind:value={value} {autofocus}/>
        </div>
        {#if isOpen}
            <DropdownMenu
                    value={currentSuggestions[currentSelectedSuggestionInList] ? currentSuggestions[currentSelectedSuggestionInList].id : ""}
                    isActiveCondition={(val1, val2) => val1 === val2}
                    data={currentSuggestions} onSelectHandler={onSelect}/>
        {/if}
    </ClickOutside>
</div>

