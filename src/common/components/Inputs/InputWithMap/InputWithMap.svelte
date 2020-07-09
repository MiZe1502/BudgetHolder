<script lang="ts">
    import InputWithIcon from "../InputWithIcon/InputWithIcon.svelte";
    import MapActionElement
        from "../../MapActionElement/MapActionElement.svelte";
    import {MapItemData} from "../../MapActionElement/utils";

    export let value = "";
    export let autofocus = false;
    export let name = "";
    export let type = "text";
    export let label = "";
    export let disabled = false;
    export let required = false;
    export let data: MapItemData = null;

    const initialAddress = data && data.address;

    const updateAddress = (newAddress: string) => {
        console.log(newAddress);
        data.address = newAddress;
        console.log(data)
    }

    const onCancel = () => {
        data.address = initialAddress;
    }

    export let onChange = () => {
    };
</script>

<InputWithIcon onChange={onChange} bind:value={value} {autofocus} {name} {type}
               {label} {disabled} {required}>
    <MapActionElement data={[data]} isEditable={true}
                      updateAddress={updateAddress}
                      onCancel={onCancel}/>
</InputWithIcon>