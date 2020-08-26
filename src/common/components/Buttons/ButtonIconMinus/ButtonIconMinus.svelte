<script lang="typescript">
    import ButtonIcon from "../ButtonIcon/ButtonIcon.svelte";
    import {style} from "./style";

    export let width: number = style.defaultButtonIconSize;
    export let height: number = style.defaultButtonIconSize;
    export let onClickHandler = () => {
    };
</script>

<ButtonIcon className={style.Button} negative={true}
            onClickHandler={onClickHandler}>
    <svg version="1.1" id="Capa_1" xmlns="http://www.w3.org/2000/svg"
         xmlns:xlink="http://www.w3.org/1999/xlink" x="0px" y="0px"
         viewBox="0 0 409.6 409.6"
         width={width} height={height}
         style="enable-background:new 0 0 409.6 409.6;" xml:space="preserve">
        <g>
            <g>
                <path d="M392.533,187.733H17.067C7.641,187.733,0,195.374,0,204.8s7.641,17.067,17.067,17.067h375.467
                    c9.426,0,17.067-7.641,17.067-17.067S401.959,187.733,392.533,187.733z"/>
            </g>
        </g>
    </svg>
</ButtonIcon>



