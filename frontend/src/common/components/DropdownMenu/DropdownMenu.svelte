<script lang="ts">
    import SimpleBar from '@woden/svelte-simplebar';
    import {DropdownData} from "./utils";
    import {style} from "./style";
    import {onMount} from "svelte";


    onMount(() => {
        console.log("DATA", data)
    });

    export let value: string | number = "";
    export let isActiveCondition: (firstValue, secondValue) => boolean;
    export let data: DropdownData[] = [];
    export let onSelectHandler: (item: DropdownData) => void;
</script>

<div class={style.Dropdown}>
    <SimpleBar style="max-height: {style.maxHeight}px; width: 100%">
        {#each data as item (item.id)}
            <div on:click={() => onSelectHandler(item)}
                 class="{isActiveCondition(item.id, value) && style.SelectedLine} {style.SingleLine} {style.FlexHorCenter} {style.Font312Black}">{item.value}</div>
        {/each}
    </SimpleBar>
</div>