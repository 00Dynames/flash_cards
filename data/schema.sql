// Download language data from tatoeba project

CREATE TABLE sentences (
"sentence_id" int not null,
"language" text not null,
"text" text not null
);

CREATE TABLE links (
"sentence_id" int not null,
"translation_id" int not null
);

CREATE TABLE eng_sentences (
"sentence_id" int not null,
"language" text not null,
"text" text not null
);

CREATE TABLE jpn_sentences (
"sentence_id" int not null,
"language" text not null,
"text" text not null
);

CREATE TABLE transcriptions (
"sentence_id" int not null,
"language" text not null,
"script_name" text not null,
"username" text not null,
"transcription" text not null
);

CREATE VIEW jpn_eng_phrases (jpn_sentence_id, jpn_text, jpn_transcription, eng_sentence_id, eng_text) as
select jpn_sentences.sentence_id, jpn_sentences.text, transcriptions.transcription, eng_sentences.sentence_id, eng_sentences.text
from jpn_sentences 
    inner join links on links.sentence_id = jpn_sentences.sentence_id
    inner join eng_sentences on links.translation_id = eng_sentences.sentence_id
    left join transcriptions on jpn_sentences.sentence_id = transcriptions.sentence_id
/* jpn_eng_phrases(jpn_sentence_id,jpn_text,jpn_transcription,eng_sentence_id,eng_text) */;