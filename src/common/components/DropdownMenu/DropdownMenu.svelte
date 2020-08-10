<script lang="ts">
    import SimpleBar from '@woden/svelte-simplebar';
    import {DropdownData} from "./utils";
    import {
        Dropdown,
        maxHeight,
        SelectedLine,
        SingleLine,
        FlexHorCenter,
        Font312Black
    } from "./style";

    export let value: string | number = "";
    export let isActiveCondition: (firstValue, secondValue) => boolean;
    export let data: DropdownData[] = [];
    export let onSelectHandler: (id: number) => void;
</script>

<div class={Dropdown}>
    <SimpleBar style="max-height: {maxHeight}px; width: 100%">
        {#each data as item (item.id)}
            <div on:click={() => onSelectHandler(item.id)}
                 class="{isActiveCondition(item.id, value) && SelectedLine} {SingleLine} {FlexHorCenter} {Font312Black}">{item.value}</div>
        {/each}
    </SimpleBar>
</div>